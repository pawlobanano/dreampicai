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

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, auth.Login())
}

func HandleSignupIndex(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, auth.Signup())
}

func HandleLoginCreate(s types.Server, w http.ResponseWriter, r *http.Request) error {
	credentials := supabase.UserCredentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	resp, err := sb.Client.Auth.SignIn(r.Context(), credentials)
	if err != nil {
		s.Logger.Error(r.Context(), "login error", "err", err)
		return render(w, r, auth.LoginForm(credentials, auth.LoginErrors{
			InvalidCredentials: "The credentials you have entered are invalid",
		}))
	}
	if err := setAuthSession(s, w, r, resp.AccessToken); err != nil {
		return err
	}
	return hxRedirect(w, r, "/")
}

func HandleLogoutCreate(s types.Server, w http.ResponseWriter, r *http.Request) error {
	store := sessions.NewCookieStore([]byte(s.Config.SessionSecret))
	session, err := store.Get(r, s.Config.SessionUserKey)
	if err != nil {
		return err
	}
	session.Values[s.Config.SessionAccessTokenKey] = ""
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return nil
}

func HandleLoginWithGoogle(w http.ResponseWriter, r *http.Request) error {
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

func HandleSignupCreate(w http.ResponseWriter, r *http.Request) error {
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

func HandleAuthCallback(s types.Server, w http.ResponseWriter, r *http.Request) error {
	accessToken := r.URL.Query().Get("access_token")
	if len(accessToken) == 0 {
		return render(w, r, auth.CallbackScript())
	}
	if err := setAuthSession(s, w, r, accessToken); err != nil {
		return err
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func setAuthSession(s types.Server, w http.ResponseWriter, r *http.Request, accessToken string) error {
	store := sessions.NewCookieStore([]byte(s.Config.SessionSecret))
	session, err := store.Get(r, s.Config.SessionUserKey)
	if err != nil {
		return err
	}
	session.Values[s.Config.SessionAccessTokenKey] = accessToken
	return session.Save(r, w)
}
