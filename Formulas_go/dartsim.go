package main

// Dartsim - Dynamic Animation and Robotics Toolkit
// Homepage: https://dartsim.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installDartsim() {
	// Método 1: Descargar y extraer .tar.gz
	dartsim_tar_url := "https://github.com/dartsim/dart/archive/refs/tags/v6.14.4.tar.gz"
	dartsim_cmd_tar := exec.Command("curl", "-L", dartsim_tar_url, "-o", "package.tar.gz")
	err := dartsim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dartsim_zip_url := "https://github.com/dartsim/dart/archive/refs/tags/v6.14.4.zip"
	dartsim_cmd_zip := exec.Command("curl", "-L", dartsim_zip_url, "-o", "package.zip")
	err = dartsim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dartsim_bin_url := "https://github.com/dartsim/dart/archive/refs/tags/v6.14.4.bin"
	dartsim_cmd_bin := exec.Command("curl", "-L", dartsim_bin_url, "-o", "binary.bin")
	err = dartsim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dartsim_src_url := "https://github.com/dartsim/dart/archive/refs/tags/v6.14.4.src.tar.gz"
	dartsim_cmd_src := exec.Command("curl", "-L", dartsim_src_url, "-o", "source.tar.gz")
	err = dartsim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dartsim_cmd_direct := exec.Command("./binary")
	err = dartsim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: assimp")
exec.Command("latte", "install", "assimp")
	fmt.Println("Instalando dependencia: bullet")
exec.Command("latte", "install", "bullet")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: fcl")
exec.Command("latte", "install", "fcl")
	fmt.Println("Instalando dependencia: flann")
exec.Command("latte", "install", "flann")
	fmt.Println("Instalando dependencia: fmt")
exec.Command("latte", "install", "fmt")
	fmt.Println("Instalando dependencia: ipopt")
exec.Command("latte", "install", "ipopt")
	fmt.Println("Instalando dependencia: libccd")
exec.Command("latte", "install", "libccd")
	fmt.Println("Instalando dependencia: nlopt")
exec.Command("latte", "install", "nlopt")
	fmt.Println("Instalando dependencia: octomap")
exec.Command("latte", "install", "octomap")
	fmt.Println("Instalando dependencia: ode")
exec.Command("latte", "install", "ode")
	fmt.Println("Instalando dependencia: open-scene-graph")
exec.Command("latte", "install", "open-scene-graph")
	fmt.Println("Instalando dependencia: spdlog")
exec.Command("latte", "install", "spdlog")
	fmt.Println("Instalando dependencia: tinyxml2")
exec.Command("latte", "install", "tinyxml2")
	fmt.Println("Instalando dependencia: urdfdom")
exec.Command("latte", "install", "urdfdom")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
