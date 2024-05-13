package handler

import (
	"dreampicai/pkg/sb"
	"dreampicai/pkg/validate"
	"dreampicai/types"
	"dreampicai/view/auth"
	"net/http"

	"github.com/nedpals/supabase-go"
)

const cookieName = "at"

func HandleLoginIndex(log types.Logger, w http.ResponseWriter, r *http.Request) error {
	return render(w, r, auth.Login())
}

func HandleSignupIndex(log types.Logger, w http.ResponseWriter, r *http.Request) error {
	return render(w, r, auth.Signup())
}

func HandleLoginCreate(log types.Logger, w http.ResponseWriter, r *http.Request) error {
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
	setAuthCookie(w, resp.AccessToken)
	return hxRedirect(w, r, "/")
}

func HandleLogoutCreate(log types.Logger, w http.ResponseWriter, r *http.Request) error {
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return nil
}

func HandleLoginWithGoogle(log types.Logger, w http.ResponseWriter, r *http.Request) error {
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

func HandleSignupCreate(log types.Logger, w http.ResponseWriter, r *http.Request) error {
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

func HandleAuthCallback(log types.Logger, w http.ResponseWriter, r *http.Request) error {
	accessToken := r.URL.Query().Get("access_token")
	if len(accessToken) == 0 {
		if err := render(w, r, auth.CallbackScript()); err != nil {
			return err
		}
	}
	setAuthCookie(w, accessToken)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func setAuthCookie(w http.ResponseWriter, accessToken string) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)
}
