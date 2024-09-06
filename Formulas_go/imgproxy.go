package main

// Imgproxy - Fast and secure server for resizing and converting remote images
// Homepage: https://imgproxy.net

import (
	"fmt"
	
	"os/exec"
)

func installImgproxy() {
	// Método 1: Descargar y extraer .tar.gz
	imgproxy_tar_url := "https://github.com/imgproxy/imgproxy/archive/refs/tags/v3.25.0.tar.gz"
	imgproxy_cmd_tar := exec.Command("curl", "-L", imgproxy_tar_url, "-o", "package.tar.gz")
	err := imgproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imgproxy_zip_url := "https://github.com/imgproxy/imgproxy/archive/refs/tags/v3.25.0.zip"
	imgproxy_cmd_zip := exec.Command("curl", "-L", imgproxy_zip_url, "-o", "package.zip")
	err = imgproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imgproxy_bin_url := "https://github.com/imgproxy/imgproxy/archive/refs/tags/v3.25.0.bin"
	imgproxy_cmd_bin := exec.Command("curl", "-L", imgproxy_bin_url, "-o", "binary.bin")
	err = imgproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imgproxy_src_url := "https://github.com/imgproxy/imgproxy/archive/refs/tags/v3.25.0.src.tar.gz"
	imgproxy_cmd_src := exec.Command("curl", "-L", imgproxy_src_url, "-o", "source.tar.gz")
	err = imgproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imgproxy_cmd_direct := exec.Command("./binary")
	err = imgproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: vips")
exec.Command("latte", "install", "vips")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
