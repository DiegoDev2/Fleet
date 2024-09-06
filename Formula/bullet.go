package main

// Bullet - Physics SDK
// Homepage: https://bulletphysics.org/

import (
	"fmt"
	
	"os/exec"
)

func installBullet() {
	// Método 1: Descargar y extraer .tar.gz
	bullet_tar_url := "https://github.com/bulletphysics/bullet3/archive/refs/tags/3.25.tar.gz"
	bullet_cmd_tar := exec.Command("curl", "-L", bullet_tar_url, "-o", "package.tar.gz")
	err := bullet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bullet_zip_url := "https://github.com/bulletphysics/bullet3/archive/refs/tags/3.25.zip"
	bullet_cmd_zip := exec.Command("curl", "-L", bullet_zip_url, "-o", "package.zip")
	err = bullet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bullet_bin_url := "https://github.com/bulletphysics/bullet3/archive/refs/tags/3.25.bin"
	bullet_cmd_bin := exec.Command("curl", "-L", bullet_bin_url, "-o", "binary.bin")
	err = bullet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bullet_src_url := "https://github.com/bulletphysics/bullet3/archive/refs/tags/3.25.src.tar.gz"
	bullet_cmd_src := exec.Command("curl", "-L", bullet_src_url, "-o", "source.tar.gz")
	err = bullet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bullet_cmd_direct := exec.Command("./binary")
	err = bullet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
