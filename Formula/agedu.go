package main

// Agedu - Unix utility for tracking down wasted disk space
// Homepage: https://www.chiark.greenend.org.uk/~sgtatham/agedu/

import (
	"fmt"
	
	"os/exec"
)

func installAgedu() {
	// Método 1: Descargar y extraer .tar.gz
	agedu_tar_url := "https://www.chiark.greenend.org.uk/~sgtatham/agedu/agedu-20211129.8cd63c5.tar.gz"
	agedu_cmd_tar := exec.Command("curl", "-L", agedu_tar_url, "-o", "package.tar.gz")
	err := agedu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	agedu_zip_url := "https://www.chiark.greenend.org.uk/~sgtatham/agedu/agedu-20211129.8cd63c5.zip"
	agedu_cmd_zip := exec.Command("curl", "-L", agedu_zip_url, "-o", "package.zip")
	err = agedu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	agedu_bin_url := "https://www.chiark.greenend.org.uk/~sgtatham/agedu/agedu-20211129.8cd63c5.bin"
	agedu_cmd_bin := exec.Command("curl", "-L", agedu_bin_url, "-o", "binary.bin")
	err = agedu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	agedu_src_url := "https://www.chiark.greenend.org.uk/~sgtatham/agedu/agedu-20211129.8cd63c5.src.tar.gz"
	agedu_cmd_src := exec.Command("curl", "-L", agedu_src_url, "-o", "source.tar.gz")
	err = agedu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	agedu_cmd_direct := exec.Command("./binary")
	err = agedu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
