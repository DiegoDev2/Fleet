package main

// Pinfo - User-friendly, console-based viewer for Info documents
// Homepage: https://packages.debian.org/sid/pinfo

import (
	"fmt"
	
	"os/exec"
)

func installPinfo() {
	// Método 1: Descargar y extraer .tar.gz
	pinfo_tar_url := "https://github.com/baszoetekouw/pinfo/archive/refs/tags/v0.6.13.tar.gz"
	pinfo_cmd_tar := exec.Command("curl", "-L", pinfo_tar_url, "-o", "package.tar.gz")
	err := pinfo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pinfo_zip_url := "https://github.com/baszoetekouw/pinfo/archive/refs/tags/v0.6.13.zip"
	pinfo_cmd_zip := exec.Command("curl", "-L", pinfo_zip_url, "-o", "package.zip")
	err = pinfo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pinfo_bin_url := "https://github.com/baszoetekouw/pinfo/archive/refs/tags/v0.6.13.bin"
	pinfo_cmd_bin := exec.Command("curl", "-L", pinfo_bin_url, "-o", "binary.bin")
	err = pinfo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pinfo_src_url := "https://github.com/baszoetekouw/pinfo/archive/refs/tags/v0.6.13.src.tar.gz"
	pinfo_cmd_src := exec.Command("curl", "-L", pinfo_src_url, "-o", "source.tar.gz")
	err = pinfo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pinfo_cmd_direct := exec.Command("./binary")
	err = pinfo_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
}
