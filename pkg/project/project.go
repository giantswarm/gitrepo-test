package project

var (
	description        = "Repository used in gitrepo tests."
	gitSHA             = "n/a"
	name        string = "gitrepo-test"
	source      string = "https://github.com/giantswarm/gitrepo-test"
	version            = "0.1.0"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
