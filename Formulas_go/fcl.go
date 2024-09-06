package main

// Fcl - Flexible Collision Library
// Homepage: https://flexible-collision-library.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installFcl() {
	// Método 1: Descargar y extraer .tar.gz
	fcl_tar_url := "https://github.com/flexible-collision-library/fcl/archive/refs/tags/0.7.0.tar.gz"
	fcl_cmd_tar := exec.Command("curl", "-L", fcl_tar_url, "-o", "package.tar.gz")
	err := fcl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fcl_zip_url := "https://github.com/flexible-collision-library/fcl/archive/refs/tags/0.7.0.zip"
	fcl_cmd_zip := exec.Command("curl", "-L", fcl_zip_url, "-o", "package.zip")
	err = fcl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fcl_bin_url := "https://github.com/flexible-collision-library/fcl/archive/refs/tags/0.7.0.bin"
	fcl_cmd_bin := exec.Command("curl", "-L", fcl_bin_url, "-o", "binary.bin")
	err = fcl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fcl_src_url := "https://github.com/flexible-collision-library/fcl/archive/refs/tags/0.7.0.src.tar.gz"
	fcl_cmd_src := exec.Command("curl", "-L", fcl_src_url, "-o", "source.tar.gz")
	err = fcl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fcl_cmd_direct := exec.Command("./binary")
	err = fcl_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libccd")
exec.Command("latte", "install", "libccd")
	fmt.Println("Instalando dependencia: octomap")
exec.Command("latte", "install", "octomap")
}
