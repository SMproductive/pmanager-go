(define-module (nordic-channel packages pmanager)
  #:use-module (gnu packages gl)
  #:use-module (gnu packages pkg-config)
  #:use-module (gnu packages xorg)
  #:use-module (guix build-system go)
  #:use-module (guix git-download)
  #:use-module ((guix licenses) #:prefix license:)
  #:use-module (guix packages))
;(define-public pmanager-go
  (package
   (name "pmanager-go")
   (version "0.1");(git-version "0.1" revision commit))
   (source (origin
	    (method git-fetch)
	    (uri (git-reference
		  (url "https://github.com/SMproductive/pmanager-go")
		  (commit "473765e44b105da132bae30b708faad81fce0c24")))
	    (file-name (git-file-name name version))
	    (sha256
	     (base32
	      "17cm9s9pm00hgg3vafrq9gc60cvk52cxgs6paqfwdz76ljp1b6k8"))))
   (build-system go-build-system)
   (native-inputs
    (list libx11 libxcursor libxrandr libxinerama
	  libxi pkg-config glfw))
   (arguments
    `(#:import-path "github.com/SMproductive/pmanager-go"))
   (home-page "https://github.com/SMproductive/pmanager-go")
   (synopsis "Nordic password manager using fyne2 api")
   (description "pmanager-go is simple and straight forward. Only what you press will happen!")
   (license license:gpl3+));)

