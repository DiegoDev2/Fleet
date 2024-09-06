package main

// Cgns - CFD General Notation System
// Homepage: http://cgns.org/

import (
	"fmt"
	
	"os/exec"
)

func installCgns() {
	// Método 1: Descargar y extraer .tar.gz
	cgns_tar_url := "https://github.com/CGNS/CGNS/archive/refs/tags/v4.4.0.tar.gz"
	cgns_cmd_tar := exec.Command("curl", "-L", cgns_tar_url, "-o", "package.tar.gz")
	err := cgns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cgns_zip_url := "https://github.com/CGNS/CGNS/archive/refs/tags/v4.4.0.zip"
	cgns_cmd_zip := exec.Command("curl", "-L", cgns_zip_url, "-o", "package.zip")
	err = cgns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cgns_bin_url := "https://github.com/CGNS/CGNS/archive/refs/tags/v4.4.0.bin"
	cgns_cmd_bin := exec.Command("curl", "-L", cgns_bin_url, "-o", "binary.bin")
	err = cgns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cgns_src_url := "https://github.com/CGNS/CGNS/archive/refs/tags/v4.4.0.src.tar.gz"
	cgns_cmd_src := exec.Command("curl", "-L", cgns_src_url, "-o", "source.tar.gz")
	err = cgns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cgns_cmd_direct := exec.Command("./binary")
	err = cgns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: hdf5")
exec.Command("latte", "install", "hdf5")
	fmt.Println("Instalando dependencia: libaec")
exec.Command("latte", "install", "libaec")
}
