package main

// Ekg2 - Multiplatform, multiprotocol, plugin-based instant messenger
// Homepage: https://github.com/ekg2/ekg2

import (
	"fmt"
	
	"os/exec"
)

func installEkg2() {
	// Método 1: Descargar y extraer .tar.gz
	ekg2_tar_url := "https://src.fedoraproject.org/lookaside/extras/ekg2/ekg2-0.3.1.tar.gz/68fc05b432c34622df6561eaabef5a40/ekg2-0.3.1.tar.gz"
	ekg2_cmd_tar := exec.Command("curl", "-L", ekg2_tar_url, "-o", "package.tar.gz")
	err := ekg2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ekg2_zip_url := "https://src.fedoraproject.org/lookaside/extras/ekg2/ekg2-0.3.1.zip/68fc05b432c34622df6561eaabef5a40/ekg2-0.3.1.zip"
	ekg2_cmd_zip := exec.Command("curl", "-L", ekg2_zip_url, "-o", "package.zip")
	err = ekg2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ekg2_bin_url := "https://src.fedoraproject.org/lookaside/extras/ekg2/ekg2-0.3.1.bin/68fc05b432c34622df6561eaabef5a40/ekg2-0.3.1.bin"
	ekg2_cmd_bin := exec.Command("curl", "-L", ekg2_bin_url, "-o", "binary.bin")
	err = ekg2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ekg2_src_url := "https://src.fedoraproject.org/lookaside/extras/ekg2/ekg2-0.3.1.src.tar.gz/68fc05b432c34622df6561eaabef5a40/ekg2-0.3.1.src.tar.gz"
	ekg2_cmd_src := exec.Command("curl", "-L", ekg2_src_url, "-o", "source.tar.gz")
	err = ekg2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ekg2_cmd_direct := exec.Command("./binary")
	err = ekg2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
