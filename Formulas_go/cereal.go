package main

// Cereal - C++11 library for serialization
// Homepage: https://uscilab.github.io/cereal/

import (
	"fmt"
	
	"os/exec"
)

func installCereal() {
	// Método 1: Descargar y extraer .tar.gz
	cereal_tar_url := "https://github.com/USCiLab/cereal/archive/refs/tags/v1.3.2.tar.gz"
	cereal_cmd_tar := exec.Command("curl", "-L", cereal_tar_url, "-o", "package.tar.gz")
	err := cereal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cereal_zip_url := "https://github.com/USCiLab/cereal/archive/refs/tags/v1.3.2.zip"
	cereal_cmd_zip := exec.Command("curl", "-L", cereal_zip_url, "-o", "package.zip")
	err = cereal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cereal_bin_url := "https://github.com/USCiLab/cereal/archive/refs/tags/v1.3.2.bin"
	cereal_cmd_bin := exec.Command("curl", "-L", cereal_bin_url, "-o", "binary.bin")
	err = cereal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cereal_src_url := "https://github.com/USCiLab/cereal/archive/refs/tags/v1.3.2.src.tar.gz"
	cereal_cmd_src := exec.Command("curl", "-L", cereal_src_url, "-o", "source.tar.gz")
	err = cereal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cereal_cmd_direct := exec.Command("./binary")
	err = cereal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
