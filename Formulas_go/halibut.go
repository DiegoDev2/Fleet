package main

// Halibut - Yet another free document preparation system
// Homepage: https://www.chiark.greenend.org.uk/~sgtatham/halibut/

import (
	"fmt"
	
	"os/exec"
)

func installHalibut() {
	// Método 1: Descargar y extraer .tar.gz
	halibut_tar_url := "https://www.chiark.greenend.org.uk/~sgtatham/halibut/halibut-1.3/halibut-1.3.tar.gz"
	halibut_cmd_tar := exec.Command("curl", "-L", halibut_tar_url, "-o", "package.tar.gz")
	err := halibut_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	halibut_zip_url := "https://www.chiark.greenend.org.uk/~sgtatham/halibut/halibut-1.3/halibut-1.3.zip"
	halibut_cmd_zip := exec.Command("curl", "-L", halibut_zip_url, "-o", "package.zip")
	err = halibut_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	halibut_bin_url := "https://www.chiark.greenend.org.uk/~sgtatham/halibut/halibut-1.3/halibut-1.3.bin"
	halibut_cmd_bin := exec.Command("curl", "-L", halibut_bin_url, "-o", "binary.bin")
	err = halibut_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	halibut_src_url := "https://www.chiark.greenend.org.uk/~sgtatham/halibut/halibut-1.3/halibut-1.3.src.tar.gz"
	halibut_cmd_src := exec.Command("curl", "-L", halibut_src_url, "-o", "source.tar.gz")
	err = halibut_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	halibut_cmd_direct := exec.Command("./binary")
	err = halibut_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
