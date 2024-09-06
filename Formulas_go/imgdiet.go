package main

// Imgdiet - Optimize and resize images
// Homepage: https://git.sr.ht/~jamesponddotco/imgdiet-go

import (
	"fmt"
	
	"os/exec"
)

func installImgdiet() {
	// Método 1: Descargar y extraer .tar.gz
	imgdiet_tar_url := "https://git.sr.ht/~jamesponddotco/imgdiet-go/archive/v0.1.2.tar.gz"
	imgdiet_cmd_tar := exec.Command("curl", "-L", imgdiet_tar_url, "-o", "package.tar.gz")
	err := imgdiet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imgdiet_zip_url := "https://git.sr.ht/~jamesponddotco/imgdiet-go/archive/v0.1.2.zip"
	imgdiet_cmd_zip := exec.Command("curl", "-L", imgdiet_zip_url, "-o", "package.zip")
	err = imgdiet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imgdiet_bin_url := "https://git.sr.ht/~jamesponddotco/imgdiet-go/archive/v0.1.2.bin"
	imgdiet_cmd_bin := exec.Command("curl", "-L", imgdiet_bin_url, "-o", "binary.bin")
	err = imgdiet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imgdiet_src_url := "https://git.sr.ht/~jamesponddotco/imgdiet-go/archive/v0.1.2.src.tar.gz"
	imgdiet_cmd_src := exec.Command("curl", "-L", imgdiet_src_url, "-o", "source.tar.gz")
	err = imgdiet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imgdiet_cmd_direct := exec.Command("./binary")
	err = imgdiet_cmd_direct.Run()
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
