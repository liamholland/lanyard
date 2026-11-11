package format

const HealthCheckPath string = "/health/check"

const timeBasePath string = "/time"
const TimeStartPath string = timeBasePath + "/start"
const TimeDepPath string = timeBasePath + "/departure"
const TimeHowManyPath string = timeBasePath + "/howmany"
const TimeGameIdealWhenPath string = timeBasePath + "/ideal"

const gamesBasePath string = "/games"
const GamesResolvePath string = gamesBasePath + "/resolve"

const identityBasePath string = "/identity"
const IdentityRetrievePath string = identityBasePath + "/retrieve"
const IdentityRegisterPath string = identityBasePath + "/register"

const adminBasePath string = "/admin"
const AdminOptionsPath string = adminBasePath + "/options/update"
const AdminUsersPath string = adminBasePath + "/users"
const AdminUsersResolvedPath = AdminUsersPath + "/resolved"
