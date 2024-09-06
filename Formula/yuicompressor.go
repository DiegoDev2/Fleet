package main

// Yuicompressor - Yahoo! JavaScript and CSS compressor
// Homepage: https://yui.github.io/yuicompressor/

import (
	"fmt"
	
	"os/exec"
)

func installYuicompressor() {
	// Método 1: Descargar y extraer .tar.gz
	yuicompressor_tar_url := "https://github.com/yui/yuicompressor/releases/download/v2.4.8/yuicompressor-2.4.8.zip"
	yuicompressor_cmd_tar := exec.Command("curl", "-L", yuicompressor_tar_url, "-o", "package.tar.gz")
	err := yuicompressor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yuicompressor_zip_url := "https://github.com/yui/yuicompressor/releases/download/v2.4.8/yuicompressor-2.4.8.zip"
	yuicompressor_cmd_zip := exec.Command("curl", "-L", yuicompressor_zip_url, "-o", "package.zip")
	err = yuicompressor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yuicompressor_bin_url := "https://github.com/yui/yuicompressor/releases/download/v2.4.8/yuicompressor-2.4.8.zip"
	yuicompressor_cmd_bin := exec.Command("curl", "-L", yuicompressor_bin_url, "-o", "binary.bin")
	err = yuicompressor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yuicompressor_src_url := "https://github.com/yui/yuicompressor/releases/download/v2.4.8/yuicompressor-2.4.8.zip"
	yuicompressor_cmd_src := exec.Command("curl", "-L", yuicompressor_src_url, "-o", "source.tar.gz")
	err = yuicompressor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yuicompressor_cmd_direct := exec.Command("./binary")
	err = yuicompressor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
