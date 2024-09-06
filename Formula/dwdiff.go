package main

// Dwdiff - Diff that operates at the word level
// Homepage: https://os.ghalkes.nl/dwdiff.html

import (
	"fmt"
	
	"os/exec"
)

func installDwdiff() {
	// Método 1: Descargar y extraer .tar.gz
	dwdiff_tar_url := "https://os.ghalkes.nl/dist/dwdiff-2.1.4.tar.bz2"
	dwdiff_cmd_tar := exec.Command("curl", "-L", dwdiff_tar_url, "-o", "package.tar.gz")
	err := dwdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dwdiff_zip_url := "https://os.ghalkes.nl/dist/dwdiff-2.1.4.tar.bz2"
	dwdiff_cmd_zip := exec.Command("curl", "-L", dwdiff_zip_url, "-o", "package.zip")
	err = dwdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dwdiff_bin_url := "https://os.ghalkes.nl/dist/dwdiff-2.1.4.tar.bz2"
	dwdiff_cmd_bin := exec.Command("curl", "-L", dwdiff_bin_url, "-o", "binary.bin")
	err = dwdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dwdiff_src_url := "https://os.ghalkes.nl/dist/dwdiff-2.1.4.tar.bz2"
	dwdiff_cmd_src := exec.Command("curl", "-L", dwdiff_src_url, "-o", "source.tar.gz")
	err = dwdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dwdiff_cmd_direct := exec.Command("./binary")
	err = dwdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
