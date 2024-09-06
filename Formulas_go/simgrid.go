package main

// Simgrid - Studies behavior of large-scale distributed systems
// Homepage: https://simgrid.org/

import (
	"fmt"
	
	"os/exec"
)

func installSimgrid() {
	// Método 1: Descargar y extraer .tar.gz
	simgrid_tar_url := "https://gitlab.inria.fr/simgrid/simgrid/-/archive/v3.35/simgrid-v3.35.tar.bz2"
	simgrid_cmd_tar := exec.Command("curl", "-L", simgrid_tar_url, "-o", "package.tar.gz")
	err := simgrid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	simgrid_zip_url := "https://gitlab.inria.fr/simgrid/simgrid/-/archive/v3.35/simgrid-v3.35.tar.bz2"
	simgrid_cmd_zip := exec.Command("curl", "-L", simgrid_zip_url, "-o", "package.zip")
	err = simgrid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	simgrid_bin_url := "https://gitlab.inria.fr/simgrid/simgrid/-/archive/v3.35/simgrid-v3.35.tar.bz2"
	simgrid_cmd_bin := exec.Command("curl", "-L", simgrid_bin_url, "-o", "binary.bin")
	err = simgrid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	simgrid_src_url := "https://gitlab.inria.fr/simgrid/simgrid/-/archive/v3.35/simgrid-v3.35.tar.bz2"
	simgrid_cmd_src := exec.Command("curl", "-L", simgrid_src_url, "-o", "source.tar.gz")
	err = simgrid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	simgrid_cmd_direct := exec.Command("./binary")
	err = simgrid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: graphviz")
exec.Command("latte", "install", "graphviz")
}
