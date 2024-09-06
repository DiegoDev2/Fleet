package main

// OpenMesh - Generic data structure to represent and manipulate polygonal meshes
// Homepage: https://www.graphics.rwth-aachen.de/software/openmesh/

import (
	"fmt"
	
	"os/exec"
)

func installOpenMesh() {
	// Método 1: Descargar y extraer .tar.gz
	openmesh_tar_url := "https://www.graphics.rwth-aachen.de/media/openmesh_static/Releases/11.0/OpenMesh-11.0.0.tar.bz2"
	openmesh_cmd_tar := exec.Command("curl", "-L", openmesh_tar_url, "-o", "package.tar.gz")
	err := openmesh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openmesh_zip_url := "https://www.graphics.rwth-aachen.de/media/openmesh_static/Releases/11.0/OpenMesh-11.0.0.tar.bz2"
	openmesh_cmd_zip := exec.Command("curl", "-L", openmesh_zip_url, "-o", "package.zip")
	err = openmesh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openmesh_bin_url := "https://www.graphics.rwth-aachen.de/media/openmesh_static/Releases/11.0/OpenMesh-11.0.0.tar.bz2"
	openmesh_cmd_bin := exec.Command("curl", "-L", openmesh_bin_url, "-o", "binary.bin")
	err = openmesh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openmesh_src_url := "https://www.graphics.rwth-aachen.de/media/openmesh_static/Releases/11.0/OpenMesh-11.0.0.tar.bz2"
	openmesh_cmd_src := exec.Command("curl", "-L", openmesh_src_url, "-o", "source.tar.gz")
	err = openmesh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openmesh_cmd_direct := exec.Command("./binary")
	err = openmesh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
