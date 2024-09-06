package main

// Spigot - Command-line streaming exact real calculator
// Homepage: https://www.chiark.greenend.org.uk/~sgtatham/spigot/

import (
	"fmt"
	
	"os/exec"
)

func installSpigot() {
	// Método 1: Descargar y extraer .tar.gz
	spigot_tar_url := "https://www.chiark.greenend.org.uk/~sgtatham/spigot/spigot-20240202.10bd0ca.tar.gz"
	spigot_cmd_tar := exec.Command("curl", "-L", spigot_tar_url, "-o", "package.tar.gz")
	err := spigot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spigot_zip_url := "https://www.chiark.greenend.org.uk/~sgtatham/spigot/spigot-20240202.10bd0ca.zip"
	spigot_cmd_zip := exec.Command("curl", "-L", spigot_zip_url, "-o", "package.zip")
	err = spigot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spigot_bin_url := "https://www.chiark.greenend.org.uk/~sgtatham/spigot/spigot-20240202.10bd0ca.bin"
	spigot_cmd_bin := exec.Command("curl", "-L", spigot_bin_url, "-o", "binary.bin")
	err = spigot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spigot_src_url := "https://www.chiark.greenend.org.uk/~sgtatham/spigot/spigot-20240202.10bd0ca.src.tar.gz"
	spigot_cmd_src := exec.Command("curl", "-L", spigot_src_url, "-o", "source.tar.gz")
	err = spigot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spigot_cmd_direct := exec.Command("./binary")
	err = spigot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
}
