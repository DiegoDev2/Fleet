package main

// RangeV3 - Experimental range library for C++14/17/20
// Homepage: https://ericniebler.github.io/range-v3/

import (
	"fmt"
	
	"os/exec"
)

func installRangeV3() {
	// Método 1: Descargar y extraer .tar.gz
	rangev3_tar_url := "https://github.com/ericniebler/range-v3/archive/refs/tags/0.12.0.tar.gz"
	rangev3_cmd_tar := exec.Command("curl", "-L", rangev3_tar_url, "-o", "package.tar.gz")
	err := rangev3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rangev3_zip_url := "https://github.com/ericniebler/range-v3/archive/refs/tags/0.12.0.zip"
	rangev3_cmd_zip := exec.Command("curl", "-L", rangev3_zip_url, "-o", "package.zip")
	err = rangev3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rangev3_bin_url := "https://github.com/ericniebler/range-v3/archive/refs/tags/0.12.0.bin"
	rangev3_cmd_bin := exec.Command("curl", "-L", rangev3_bin_url, "-o", "binary.bin")
	err = rangev3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rangev3_src_url := "https://github.com/ericniebler/range-v3/archive/refs/tags/0.12.0.src.tar.gz"
	rangev3_cmd_src := exec.Command("curl", "-L", rangev3_src_url, "-o", "source.tar.gz")
	err = rangev3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rangev3_cmd_direct := exec.Command("./binary")
	err = rangev3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
