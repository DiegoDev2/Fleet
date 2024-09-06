package main

// Beagle - Evaluate the likelihood of sequence evolution on trees
// Homepage: https://github.com/beagle-dev/beagle-lib

import (
	"fmt"
	
	"os/exec"
)

func installBeagle() {
	// Método 1: Descargar y extraer .tar.gz
	beagle_tar_url := "https://github.com/beagle-dev/beagle-lib/archive/refs/tags/v4.0.1.tar.gz"
	beagle_cmd_tar := exec.Command("curl", "-L", beagle_tar_url, "-o", "package.tar.gz")
	err := beagle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	beagle_zip_url := "https://github.com/beagle-dev/beagle-lib/archive/refs/tags/v4.0.1.zip"
	beagle_cmd_zip := exec.Command("curl", "-L", beagle_zip_url, "-o", "package.zip")
	err = beagle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	beagle_bin_url := "https://github.com/beagle-dev/beagle-lib/archive/refs/tags/v4.0.1.bin"
	beagle_cmd_bin := exec.Command("curl", "-L", beagle_bin_url, "-o", "binary.bin")
	err = beagle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	beagle_src_url := "https://github.com/beagle-dev/beagle-lib/archive/refs/tags/v4.0.1.src.tar.gz"
	beagle_cmd_src := exec.Command("curl", "-L", beagle_src_url, "-o", "source.tar.gz")
	err = beagle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	beagle_cmd_direct := exec.Command("./binary")
	err = beagle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
}
