package main

// Field3d - Library for storing voxel data on disk and in memory
// Homepage: https://sites.google.com/site/field3d/

import (
	"fmt"
	
	"os/exec"
)

func installField3d() {
	// Método 1: Descargar y extraer .tar.gz
	field3d_tar_url := "https://github.com/imageworks/Field3D/archive/refs/tags/v1.7.3.tar.gz"
	field3d_cmd_tar := exec.Command("curl", "-L", field3d_tar_url, "-o", "package.tar.gz")
	err := field3d_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	field3d_zip_url := "https://github.com/imageworks/Field3D/archive/refs/tags/v1.7.3.zip"
	field3d_cmd_zip := exec.Command("curl", "-L", field3d_zip_url, "-o", "package.zip")
	err = field3d_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	field3d_bin_url := "https://github.com/imageworks/Field3D/archive/refs/tags/v1.7.3.bin"
	field3d_cmd_bin := exec.Command("curl", "-L", field3d_bin_url, "-o", "binary.bin")
	err = field3d_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	field3d_src_url := "https://github.com/imageworks/Field3D/archive/refs/tags/v1.7.3.src.tar.gz"
	field3d_cmd_src := exec.Command("curl", "-L", field3d_src_url, "-o", "source.tar.gz")
	err = field3d_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	field3d_cmd_direct := exec.Command("./binary")
	err = field3d_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: hdf5")
	exec.Command("latte", "install", "hdf5").Run()
	fmt.Println("Instalando dependencia: ilmbase")
	exec.Command("latte", "install", "ilmbase").Run()
}
