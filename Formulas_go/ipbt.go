package main

// Ipbt - Program for recording a UNIX terminal session
// Homepage: https://www.chiark.greenend.org.uk/~sgtatham/ipbt/

import (
	"fmt"
	
	"os/exec"
)

func installIpbt() {
	// Método 1: Descargar y extraer .tar.gz
	ipbt_tar_url := "https://www.chiark.greenend.org.uk/~sgtatham/ipbt/ipbt-20240501.bc876ea.tar.gz"
	ipbt_cmd_tar := exec.Command("curl", "-L", ipbt_tar_url, "-o", "package.tar.gz")
	err := ipbt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ipbt_zip_url := "https://www.chiark.greenend.org.uk/~sgtatham/ipbt/ipbt-20240501.bc876ea.zip"
	ipbt_cmd_zip := exec.Command("curl", "-L", ipbt_zip_url, "-o", "package.zip")
	err = ipbt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ipbt_bin_url := "https://www.chiark.greenend.org.uk/~sgtatham/ipbt/ipbt-20240501.bc876ea.bin"
	ipbt_cmd_bin := exec.Command("curl", "-L", ipbt_bin_url, "-o", "binary.bin")
	err = ipbt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ipbt_src_url := "https://www.chiark.greenend.org.uk/~sgtatham/ipbt/ipbt-20240501.bc876ea.src.tar.gz"
	ipbt_cmd_src := exec.Command("curl", "-L", ipbt_src_url, "-o", "source.tar.gz")
	err = ipbt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ipbt_cmd_direct := exec.Command("./binary")
	err = ipbt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
