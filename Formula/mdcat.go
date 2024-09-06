package main

// Mdcat - Show markdown documents on text terminals
// Homepage: https://github.com/swsnr/mdcat

import (
	"fmt"
	
	"os/exec"
)

func installMdcat() {
	// Método 1: Descargar y extraer .tar.gz
	mdcat_tar_url := "https://github.com/swsnr/mdcat/archive/refs/tags/mdcat-2.1.2.tar.gz"
	mdcat_cmd_tar := exec.Command("curl", "-L", mdcat_tar_url, "-o", "package.tar.gz")
	err := mdcat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mdcat_zip_url := "https://github.com/swsnr/mdcat/archive/refs/tags/mdcat-2.1.2.zip"
	mdcat_cmd_zip := exec.Command("curl", "-L", mdcat_zip_url, "-o", "package.zip")
	err = mdcat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mdcat_bin_url := "https://github.com/swsnr/mdcat/archive/refs/tags/mdcat-2.1.2.bin"
	mdcat_cmd_bin := exec.Command("curl", "-L", mdcat_bin_url, "-o", "binary.bin")
	err = mdcat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mdcat_src_url := "https://github.com/swsnr/mdcat/archive/refs/tags/mdcat-2.1.2.src.tar.gz"
	mdcat_cmd_src := exec.Command("curl", "-L", mdcat_src_url, "-o", "source.tar.gz")
	err = mdcat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mdcat_cmd_direct := exec.Command("./binary")
	err = mdcat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoctor")
	exec.Command("latte", "install", "asciidoctor").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
