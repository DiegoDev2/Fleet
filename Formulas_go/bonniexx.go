package main

// Bonniexx - Benchmark suite for file systems and hard drives
// Homepage: https://www.coker.com.au/bonnie++/

import (
	"fmt"
	
	"os/exec"
)

func installBonniexx() {
	// Método 1: Descargar y extraer .tar.gz
	bonniexx_tar_url := "https://www.coker.com.au/bonnie++/bonnie++-2.00a.tgz"
	bonniexx_cmd_tar := exec.Command("curl", "-L", bonniexx_tar_url, "-o", "package.tar.gz")
	err := bonniexx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bonniexx_zip_url := "https://www.coker.com.au/bonnie++/bonnie++-2.00a.tgz"
	bonniexx_cmd_zip := exec.Command("curl", "-L", bonniexx_zip_url, "-o", "package.zip")
	err = bonniexx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bonniexx_bin_url := "https://www.coker.com.au/bonnie++/bonnie++-2.00a.tgz"
	bonniexx_cmd_bin := exec.Command("curl", "-L", bonniexx_bin_url, "-o", "binary.bin")
	err = bonniexx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bonniexx_src_url := "https://www.coker.com.au/bonnie++/bonnie++-2.00a.tgz"
	bonniexx_cmd_src := exec.Command("curl", "-L", bonniexx_src_url, "-o", "source.tar.gz")
	err = bonniexx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bonniexx_cmd_direct := exec.Command("./binary")
	err = bonniexx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
