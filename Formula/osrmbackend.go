package main

// OsrmBackend - High performance routing engine
// Homepage: https://project-osrm.org/

import (
	"fmt"
	
	"os/exec"
)

func installOsrmBackend() {
	// Método 1: Descargar y extraer .tar.gz
	osrmbackend_tar_url := "https://github.com/Project-OSRM/osrm-backend/archive/refs/tags/v5.27.1.tar.gz"
	osrmbackend_cmd_tar := exec.Command("curl", "-L", osrmbackend_tar_url, "-o", "package.tar.gz")
	err := osrmbackend_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osrmbackend_zip_url := "https://github.com/Project-OSRM/osrm-backend/archive/refs/tags/v5.27.1.zip"
	osrmbackend_cmd_zip := exec.Command("curl", "-L", osrmbackend_zip_url, "-o", "package.zip")
	err = osrmbackend_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osrmbackend_bin_url := "https://github.com/Project-OSRM/osrm-backend/archive/refs/tags/v5.27.1.bin"
	osrmbackend_cmd_bin := exec.Command("curl", "-L", osrmbackend_bin_url, "-o", "binary.bin")
	err = osrmbackend_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osrmbackend_src_url := "https://github.com/Project-OSRM/osrm-backend/archive/refs/tags/v5.27.1.src.tar.gz"
	osrmbackend_cmd_src := exec.Command("curl", "-L", osrmbackend_src_url, "-o", "source.tar.gz")
	err = osrmbackend_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osrmbackend_cmd_direct := exec.Command("./binary")
	err = osrmbackend_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: libstxxl")
	exec.Command("latte", "install", "libstxxl").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: libzip")
	exec.Command("latte", "install", "libzip").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: tbb")
	exec.Command("latte", "install", "tbb").Run()
}
