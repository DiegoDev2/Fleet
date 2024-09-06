package main

// SigrokCli - Sigrok command-line interface to use logic analyzers and more
// Homepage: https://sigrok.org/

import (
	"fmt"
	
	"os/exec"
)

func installSigrokCli() {
	// Método 1: Descargar y extraer .tar.gz
	sigrokcli_tar_url := "https://sigrok.org/download/source/sigrok-cli/sigrok-cli-0.7.2.tar.gz"
	sigrokcli_cmd_tar := exec.Command("curl", "-L", sigrokcli_tar_url, "-o", "package.tar.gz")
	err := sigrokcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sigrokcli_zip_url := "https://sigrok.org/download/source/sigrok-cli/sigrok-cli-0.7.2.zip"
	sigrokcli_cmd_zip := exec.Command("curl", "-L", sigrokcli_zip_url, "-o", "package.zip")
	err = sigrokcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sigrokcli_bin_url := "https://sigrok.org/download/source/sigrok-cli/sigrok-cli-0.7.2.bin"
	sigrokcli_cmd_bin := exec.Command("curl", "-L", sigrokcli_bin_url, "-o", "binary.bin")
	err = sigrokcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sigrokcli_src_url := "https://sigrok.org/download/source/sigrok-cli/sigrok-cli-0.7.2.src.tar.gz"
	sigrokcli_cmd_src := exec.Command("curl", "-L", sigrokcli_src_url, "-o", "source.tar.gz")
	err = sigrokcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sigrokcli_cmd_direct := exec.Command("./binary")
	err = sigrokcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libsigrok")
	exec.Command("latte", "install", "libsigrok").Run()
	fmt.Println("Instalando dependencia: libsigrokdecode")
	exec.Command("latte", "install", "libsigrokdecode").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
