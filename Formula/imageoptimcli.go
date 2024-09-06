package main

// ImageoptimCli - CLI for ImageOptim, ImageAlpha and JPEGmini
// Homepage: https://jamiemason.github.io/ImageOptim-CLI/

import (
	"fmt"
	
	"os/exec"
)

func installImageoptimCli() {
	// Método 1: Descargar y extraer .tar.gz
	imageoptimcli_tar_url := "https://github.com/JamieMason/ImageOptim-CLI/archive/refs/tags/3.1.9.tar.gz"
	imageoptimcli_cmd_tar := exec.Command("curl", "-L", imageoptimcli_tar_url, "-o", "package.tar.gz")
	err := imageoptimcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imageoptimcli_zip_url := "https://github.com/JamieMason/ImageOptim-CLI/archive/refs/tags/3.1.9.zip"
	imageoptimcli_cmd_zip := exec.Command("curl", "-L", imageoptimcli_zip_url, "-o", "package.zip")
	err = imageoptimcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imageoptimcli_bin_url := "https://github.com/JamieMason/ImageOptim-CLI/archive/refs/tags/3.1.9.bin"
	imageoptimcli_cmd_bin := exec.Command("curl", "-L", imageoptimcli_bin_url, "-o", "binary.bin")
	err = imageoptimcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imageoptimcli_src_url := "https://github.com/JamieMason/ImageOptim-CLI/archive/refs/tags/3.1.9.src.tar.gz"
	imageoptimcli_cmd_src := exec.Command("curl", "-L", imageoptimcli_src_url, "-o", "source.tar.gz")
	err = imageoptimcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imageoptimcli_cmd_direct := exec.Command("./binary")
	err = imageoptimcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: yarn")
	exec.Command("latte", "install", "yarn").Run()
}
