package main

// Pidgin - Multi-protocol chat client
// Homepage: https://pidgin.im/

import (
	"fmt"
	
	"os/exec"
)

func installPidgin() {
	// Método 1: Descargar y extraer .tar.gz
	pidgin_tar_url := "https://downloads.sourceforge.net/project/pidgin/Pidgin/2.14.13/pidgin-2.14.13.tar.bz2"
	pidgin_cmd_tar := exec.Command("curl", "-L", pidgin_tar_url, "-o", "package.tar.gz")
	err := pidgin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pidgin_zip_url := "https://downloads.sourceforge.net/project/pidgin/Pidgin/2.14.13/pidgin-2.14.13.tar.bz2"
	pidgin_cmd_zip := exec.Command("curl", "-L", pidgin_zip_url, "-o", "package.zip")
	err = pidgin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pidgin_bin_url := "https://downloads.sourceforge.net/project/pidgin/Pidgin/2.14.13/pidgin-2.14.13.tar.bz2"
	pidgin_cmd_bin := exec.Command("curl", "-L", pidgin_bin_url, "-o", "binary.bin")
	err = pidgin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pidgin_src_url := "https://downloads.sourceforge.net/project/pidgin/Pidgin/2.14.13/pidgin-2.14.13.tar.bz2"
	pidgin_cmd_src := exec.Command("curl", "-L", pidgin_src_url, "-o", "source.tar.gz")
	err = pidgin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pidgin_cmd_direct := exec.Command("./binary")
	err = pidgin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: gtk+")
exec.Command("latte", "install", "gtk+")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: libgnt")
exec.Command("latte", "install", "libgnt")
	fmt.Println("Instalando dependencia: libidn")
exec.Command("latte", "install", "libidn")
	fmt.Println("Instalando dependencia: libotr")
exec.Command("latte", "install", "libotr")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: libsm")
exec.Command("latte", "install", "libsm")
	fmt.Println("Instalando dependencia: libxscrnsaver")
exec.Command("latte", "install", "libxscrnsaver")
}
