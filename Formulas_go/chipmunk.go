package main

// Chipmunk - 2D rigid body physics library written in C
// Homepage: https://chipmunk-physics.net/

import (
	"fmt"
	
	"os/exec"
)

func installChipmunk() {
	// Método 1: Descargar y extraer .tar.gz
	chipmunk_tar_url := "https://chipmunk-physics.net/release/Chipmunk-7.x/Chipmunk-7.0.3.tgz"
	chipmunk_cmd_tar := exec.Command("curl", "-L", chipmunk_tar_url, "-o", "package.tar.gz")
	err := chipmunk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chipmunk_zip_url := "https://chipmunk-physics.net/release/Chipmunk-7.x/Chipmunk-7.0.3.tgz"
	chipmunk_cmd_zip := exec.Command("curl", "-L", chipmunk_zip_url, "-o", "package.zip")
	err = chipmunk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chipmunk_bin_url := "https://chipmunk-physics.net/release/Chipmunk-7.x/Chipmunk-7.0.3.tgz"
	chipmunk_cmd_bin := exec.Command("curl", "-L", chipmunk_bin_url, "-o", "binary.bin")
	err = chipmunk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chipmunk_src_url := "https://chipmunk-physics.net/release/Chipmunk-7.x/Chipmunk-7.0.3.tgz"
	chipmunk_cmd_src := exec.Command("curl", "-L", chipmunk_src_url, "-o", "source.tar.gz")
	err = chipmunk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chipmunk_cmd_direct := exec.Command("./binary")
	err = chipmunk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
