package handler

import (
	"dreampicai/pkg/sb"
	"dreampicai/pkg/validate"
	"dreampicai/types"
	"dreampicai/view/auth"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/nedpals/supabase-go"
)

const (
	sessionAccessTokenKey = "accessToken"
	sessionUserKey        = "user"
)

func HandleLoginIndex(cfg types.Config, log types.Logger, w http.ResponseWriter, r *http.Request) error {
	return render(w, r, auth.Login())
}

func HandleSignupIndex(cfg types.Config, log types.Logger, w http.ResponseWriter, r *http.Request) error {
	return render(w, r, auth.Signup())
}

func HandleLoginCreate(cfg types.Config, log types.Logger, w http.ResponseWriter, r *http.Request) error {
	credentials := supabase.UserCredentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	resp, err := sb.Client.Auth.SignIn(r.Context(), credentials)
	if err != nil {
		log.Error(r.Context(), "login error", "err", err)
		return render(w, r, auth.LoginForm(credentials, auth.LoginErrors{
			InvalidCredentials: "The credentials you have entered are invalid",
		}))
	}
	if err := setAuthSession(cfg, w, r, resp.AccessToken); err != nil {
		return err
	}
	return hxRedirect(w, r, "/")
}

func HandleLogoutCreate(cfg types.Config, log types.Logger, w http.ResponseWriter, r *http.Request) error {
	store := sessions.NewCookieStore([]byte(cfg.SessionSecret))
	session, err := store.Get(r, sessionUserKey)
	if err != nil {
		return err
	}
	session.Values[sessionAccessTokenKey] = ""
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return nil
}

func HandleLoginWithGoogle(cfg types.Config, log types.Logger, w http.ResponseWriter, r *http.Request) error {
	resp, err := sb.Client.Auth.SignInWithProvider(supabase.ProviderSignInOptions{
		Provider:   "google",
		RedirectTo: "http://localhost:3000/auth/callback",
	})
	if err != nil {
		return err
	}
	http.Redirect(w, r, resp.URL, http.StatusSeeOther)
	return nil
}

func HandleSignupCreate(cfg types.Config, log types.Logger, w http.ResponseWriter, r *http.Request) error {
	params := auth.SignupParams{
		Email:           r.FormValue("email"),
		Password:        r.FormValue("password"),
		ConfirmPassword: r.FormValue("confirmPassword"),
	}
	errors := auth.SignupErrors{}
	if ok := validate.New(&params, validate.Fields{
		"Email":    validate.Rules(validate.Email),
		"Password": validate.Rules(validate.Password),
		"ConfirmPassword": validate.Rules(
			validate.Equal(params.Password),
			validate.Message("passwords do not match"),
		),
	}).Validate(&errors); !ok {
		return render(w, r, auth.SignupForm(params, errors))
	}
	user, err := sb.Client.Auth.SignUp(r.Context(), supabase.UserCredentials{
		Email:    params.Email,
		Password: params.Password,
	})
	if err != nil {
		return err
	}
	return render(w, r, auth.SignupSuccess(user.Email))
}

func HandleAuthCallback(cfg types.Config, log types.Logger, w http.ResponseWriter, r *http.Request) error {
	accessToken := r.URL.Query().Get("access_token")
	if len(accessToken) == 0 {
		return render(w, r, auth.CallbackScript())
	}
	if err := setAuthSession(cfg, w, r, accessToken); err != nil {
		return err
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func setAuthSession(cfg types.Config, w http.ResponseWriter, r *http.Request, accessToken string) error {
	store := sessions.NewCookieStore([]byte(cfg.SessionSecret))
	session, err := store.Get(r, sessionUserKey)
	if err != nil {
		return err
	}
	session.Values[sessionAccessTokenKey] = accessToken
	return session.Save(r, w)
}
