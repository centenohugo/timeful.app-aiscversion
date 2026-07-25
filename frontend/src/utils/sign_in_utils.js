import { calendarTypes } from "@/constants"
import store from "@/store"

/** Redirects user to the correct google sign in page */
export const signInGoogle = ({
  state = {},
  selectAccount = false,
  requestCalendarPermission = false,
  loginHint = "",
}) => {
  const clientId = process.env.VUE_APP_GOOGLE_CLIENT_ID
  const redirectUri = `${window.location.origin}/auth`

  let scope = "openid email profile "
  if (requestCalendarPermission) {
    scope +=
      "https://www.googleapis.com/auth/calendar.calendarlist.readonly https://www.googleapis.com/auth/calendar.events.readonly "
  }
  scope = encodeURIComponent(scope)

  let stateString = ""
  if (!state) state = {}
  state.calendarType = calendarTypes.GOOGLE
  state = encodeURIComponent(JSON.stringify(state))
  stateString = `&state=${state}`

  let promptString = ""
  if (selectAccount) {
    promptString = "&prompt=select_account+consent"
  } else {
    promptString = "&prompt=consent"
    if (loginHint.length > 0) {
      promptString += `&login_hint=${loginHint}`
    } else if (store.state.authUser) {
      promptString += `&login_hint=${store.state.authUser.email}`
    }
  }

  const url = `https://accounts.google.com/o/oauth2/v2/auth?client_id=${clientId}&redirect_uri=${redirectUri}&response_type=code&scope=${scope}&access_type=offline${promptString}${stateString}&include_granted_scopes=true`
  window.location.href = url
}
