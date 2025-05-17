package usecase

//go:generate mockery --disable-version-string --with-expecter --name UserRepo --filename user_repo_mock.go
//go:generate mockery --disable-version-string --with-expecter --name SessionRepo --filename session_repo_mock.go
//go:generate mockery --disable-version-string --with-expecter --name TokenService --filename token_service_mock.go
//go:generate mockery --disable-version-string --with-expecter --name SessionLister --filename session_lister_mock.go
