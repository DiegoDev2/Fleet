package main

// Profanity - Console based XMPP client
// Homepage: https://profanity-im.github.io

import (
	"fmt"
	
	"os/exec"
)

func installProfanity() {
	// Método 1: Descargar y extraer .tar.gz
	profanity_tar_url := "https://profanity-im.github.io/tarballs/profanity-0.14.0.tar.gz"
	profanity_cmd_tar := exec.Command("curl", "-L", profanity_tar_url, "-o", "package.tar.gz")
	err := profanity_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	profanity_zip_url := "https://profanity-im.github.io/tarballs/profanity-0.14.0.zip"
	profanity_cmd_zip := exec.Command("curl", "-L", profanity_zip_url, "-o", "package.zip")
	err = profanity_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	profanity_bin_url := "https://profanity-im.github.io/tarballs/profanity-0.14.0.bin"
	profanity_cmd_bin := exec.Command("curl", "-L", profanity_bin_url, "-o", "binary.bin")
	err = profanity_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	profanity_src_url := "https://profanity-im.github.io/tarballs/profanity-0.14.0.src.tar.gz"
	profanity_cmd_src := exec.Command("curl", "-L", profanity_src_url, "-o", "source.tar.gz")
	err = profanity_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	profanity_cmd_direct := exec.Command("./binary")
	err = profanity_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: autoconf-archive")
exec.Command("latte", "install", "autoconf-archive")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: libomemo-c")
exec.Command("latte", "install", "libomemo-c")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: curl")
exec.Command("latte", "install", "curl")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: gpgme")
exec.Command("latte", "install", "gpgme")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: libotr")
exec.Command("latte", "install", "libotr")
	fmt.Println("Instalando dependencia: libstrophe")
exec.Command("latte", "install", "libstrophe")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libassuan")
exec.Command("latte", "install", "libassuan")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
	fmt.Println("Instalando dependencia: terminal-notifier")
exec.Command("latte", "install", "terminal-notifier")
}
