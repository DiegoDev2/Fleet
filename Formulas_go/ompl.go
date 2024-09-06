package main

// Ompl - Open Motion Planning Library consists of many motion planning algorithms
// Homepage: https://ompl.kavrakilab.org/

import (
	"fmt"
	
	"os/exec"
)

func installOmpl() {
	// Método 1: Descargar y extraer .tar.gz
	ompl_tar_url := "https://github.com/ompl/ompl/archive/refs/tags/1.6.0.tar.gz"
	ompl_cmd_tar := exec.Command("curl", "-L", ompl_tar_url, "-o", "package.tar.gz")
	err := ompl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ompl_zip_url := "https://github.com/ompl/ompl/archive/refs/tags/1.6.0.zip"
	ompl_cmd_zip := exec.Command("curl", "-L", ompl_zip_url, "-o", "package.zip")
	err = ompl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ompl_bin_url := "https://github.com/ompl/ompl/archive/refs/tags/1.6.0.bin"
	ompl_cmd_bin := exec.Command("curl", "-L", ompl_bin_url, "-o", "binary.bin")
	err = ompl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ompl_src_url := "https://github.com/ompl/ompl/archive/refs/tags/1.6.0.src.tar.gz"
	ompl_cmd_src := exec.Command("curl", "-L", ompl_src_url, "-o", "source.tar.gz")
	err = ompl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ompl_cmd_direct := exec.Command("./binary")
	err = ompl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: flann")
exec.Command("latte", "install", "flann")
	fmt.Println("Instalando dependencia: ode")
exec.Command("latte", "install", "ode")
}
