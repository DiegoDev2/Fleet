package main

// Lgogdownloader - Unofficial downloader for GOG.com games
// Homepage: https://sites.google.com/site/gogdownloader/

import (
	"fmt"
	
	"os/exec"
)

func installLgogdownloader() {
	// Método 1: Descargar y extraer .tar.gz
	lgogdownloader_tar_url := "https://github.com/Sude-/lgogdownloader/releases/download/v3.15/lgogdownloader-3.15.tar.gz"
	lgogdownloader_cmd_tar := exec.Command("curl", "-L", lgogdownloader_tar_url, "-o", "package.tar.gz")
	err := lgogdownloader_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lgogdownloader_zip_url := "https://github.com/Sude-/lgogdownloader/releases/download/v3.15/lgogdownloader-3.15.zip"
	lgogdownloader_cmd_zip := exec.Command("curl", "-L", lgogdownloader_zip_url, "-o", "package.zip")
	err = lgogdownloader_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lgogdownloader_bin_url := "https://github.com/Sude-/lgogdownloader/releases/download/v3.15/lgogdownloader-3.15.bin"
	lgogdownloader_cmd_bin := exec.Command("curl", "-L", lgogdownloader_bin_url, "-o", "binary.bin")
	err = lgogdownloader_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lgogdownloader_src_url := "https://github.com/Sude-/lgogdownloader/releases/download/v3.15/lgogdownloader-3.15.src.tar.gz"
	lgogdownloader_cmd_src := exec.Command("curl", "-L", lgogdownloader_src_url, "-o", "source.tar.gz")
	err = lgogdownloader_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lgogdownloader_cmd_direct := exec.Command("./binary")
	err = lgogdownloader_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: help2man")
	exec.Command("latte", "install", "help2man").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: htmlcxx")
	exec.Command("latte", "install", "htmlcxx").Run()
	fmt.Println("Instalando dependencia: jsoncpp")
	exec.Command("latte", "install", "jsoncpp").Run()
	fmt.Println("Instalando dependencia: rhash")
	exec.Command("latte", "install", "rhash").Run()
	fmt.Println("Instalando dependencia: tidy-html5")
	exec.Command("latte", "install", "tidy-html5").Run()
	fmt.Println("Instalando dependencia: tinyxml2")
	exec.Command("latte", "install", "tinyxml2").Run()
}
