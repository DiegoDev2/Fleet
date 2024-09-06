package main

// Notmuch - Thread-based email index, search, and tagging
// Homepage: https://notmuchmail.org/

import (
	"fmt"
	
	"os/exec"
)

func installNotmuch() {
	// Método 1: Descargar y extraer .tar.gz
	notmuch_tar_url := "https://notmuchmail.org/releases/notmuch-0.38.3.tar.xz"
	notmuch_cmd_tar := exec.Command("curl", "-L", notmuch_tar_url, "-o", "package.tar.gz")
	err := notmuch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	notmuch_zip_url := "https://notmuchmail.org/releases/notmuch-0.38.3.tar.xz"
	notmuch_cmd_zip := exec.Command("curl", "-L", notmuch_zip_url, "-o", "package.zip")
	err = notmuch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	notmuch_bin_url := "https://notmuchmail.org/releases/notmuch-0.38.3.tar.xz"
	notmuch_cmd_bin := exec.Command("curl", "-L", notmuch_bin_url, "-o", "binary.bin")
	err = notmuch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	notmuch_src_url := "https://notmuchmail.org/releases/notmuch-0.38.3.tar.xz"
	notmuch_cmd_src := exec.Command("curl", "-L", notmuch_src_url, "-o", "source.tar.gz")
	err = notmuch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	notmuch_cmd_direct := exec.Command("./binary")
	err = notmuch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: emacs")
	exec.Command("latte", "install", "emacs").Run()
	fmt.Println("Instalando dependencia: libgpg-error")
	exec.Command("latte", "install", "libgpg-error").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
	fmt.Println("Instalando dependencia: cffi")
	exec.Command("latte", "install", "cffi").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gmime")
	exec.Command("latte", "install", "gmime").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: talloc")
	exec.Command("latte", "install", "talloc").Run()
	fmt.Println("Instalando dependencia: xapian")
	exec.Command("latte", "install", "xapian").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
