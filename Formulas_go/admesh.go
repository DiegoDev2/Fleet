package main

// Admesh - Processes triangulated solid meshes
// Homepage: https://github.com/admesh/admesh

import (
	"fmt"
	
	"os/exec"
)

func installAdmesh() {
	// Método 1: Descargar y extraer .tar.gz
	admesh_tar_url := "https://github.com/admesh/admesh/releases/download/v0.98.5/admesh-0.98.5.tar.gz"
	admesh_cmd_tar := exec.Command("curl", "-L", admesh_tar_url, "-o", "package.tar.gz")
	err := admesh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	admesh_zip_url := "https://github.com/admesh/admesh/releases/download/v0.98.5/admesh-0.98.5.zip"
	admesh_cmd_zip := exec.Command("curl", "-L", admesh_zip_url, "-o", "package.zip")
	err = admesh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	admesh_bin_url := "https://github.com/admesh/admesh/releases/download/v0.98.5/admesh-0.98.5.bin"
	admesh_cmd_bin := exec.Command("curl", "-L", admesh_bin_url, "-o", "binary.bin")
	err = admesh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	admesh_src_url := "https://github.com/admesh/admesh/releases/download/v0.98.5/admesh-0.98.5.src.tar.gz"
	admesh_cmd_src := exec.Command("curl", "-L", admesh_src_url, "-o", "source.tar.gz")
	err = admesh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	admesh_cmd_direct := exec.Command("./binary")
	err = admesh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
