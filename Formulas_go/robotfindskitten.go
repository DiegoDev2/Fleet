package main

// Robotfindskitten - Zen Simulation of robot finding kitten
// Homepage: http://robotfindskitten.org/

import (
	"fmt"
	
	"os/exec"
)

func installRobotfindskitten() {
	// Método 1: Descargar y extraer .tar.gz
	robotfindskitten_tar_url := "https://downloads.sourceforge.net/project/rfk/robotfindskitten-POSIX/ship_it_anyway/robotfindskitten-2.8284271.702.tar.gz"
	robotfindskitten_cmd_tar := exec.Command("curl", "-L", robotfindskitten_tar_url, "-o", "package.tar.gz")
	err := robotfindskitten_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	robotfindskitten_zip_url := "https://downloads.sourceforge.net/project/rfk/robotfindskitten-POSIX/ship_it_anyway/robotfindskitten-2.8284271.702.zip"
	robotfindskitten_cmd_zip := exec.Command("curl", "-L", robotfindskitten_zip_url, "-o", "package.zip")
	err = robotfindskitten_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	robotfindskitten_bin_url := "https://downloads.sourceforge.net/project/rfk/robotfindskitten-POSIX/ship_it_anyway/robotfindskitten-2.8284271.702.bin"
	robotfindskitten_cmd_bin := exec.Command("curl", "-L", robotfindskitten_bin_url, "-o", "binary.bin")
	err = robotfindskitten_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	robotfindskitten_src_url := "https://downloads.sourceforge.net/project/rfk/robotfindskitten-POSIX/ship_it_anyway/robotfindskitten-2.8284271.702.src.tar.gz"
	robotfindskitten_cmd_src := exec.Command("curl", "-L", robotfindskitten_src_url, "-o", "source.tar.gz")
	err = robotfindskitten_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	robotfindskitten_cmd_direct := exec.Command("./binary")
	err = robotfindskitten_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
