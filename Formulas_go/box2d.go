package main

// Box2d - 2D physics engine for games
// Homepage: https://box2d.org

import (
	"fmt"
	
	"os/exec"
)

func installBox2d() {
	// Método 1: Descargar y extraer .tar.gz
	box2d_tar_url := "https://github.com/erincatto/box2d/archive/refs/tags/v2.4.2.tar.gz"
	box2d_cmd_tar := exec.Command("curl", "-L", box2d_tar_url, "-o", "package.tar.gz")
	err := box2d_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	box2d_zip_url := "https://github.com/erincatto/box2d/archive/refs/tags/v2.4.2.zip"
	box2d_cmd_zip := exec.Command("curl", "-L", box2d_zip_url, "-o", "package.zip")
	err = box2d_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	box2d_bin_url := "https://github.com/erincatto/box2d/archive/refs/tags/v2.4.2.bin"
	box2d_cmd_bin := exec.Command("curl", "-L", box2d_bin_url, "-o", "binary.bin")
	err = box2d_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	box2d_src_url := "https://github.com/erincatto/box2d/archive/refs/tags/v2.4.2.src.tar.gz"
	box2d_cmd_src := exec.Command("curl", "-L", box2d_src_url, "-o", "source.tar.gz")
	err = box2d_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	box2d_cmd_direct := exec.Command("./binary")
	err = box2d_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doctest")
exec.Command("latte", "install", "doctest")
}
