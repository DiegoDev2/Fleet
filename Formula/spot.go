package main

// Spot - Platform for LTL and ω-automata manipulation
// Homepage: https://spot.lre.epita.fr

import (
	"fmt"
	
	"os/exec"
)

func installSpot() {
	// Método 1: Descargar y extraer .tar.gz
	spot_tar_url := "https://www.lrde.epita.fr/dload/spot/spot-2.12.tar.gz"
	spot_cmd_tar := exec.Command("curl", "-L", spot_tar_url, "-o", "package.tar.gz")
	err := spot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spot_zip_url := "https://www.lrde.epita.fr/dload/spot/spot-2.12.zip"
	spot_cmd_zip := exec.Command("curl", "-L", spot_zip_url, "-o", "package.zip")
	err = spot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spot_bin_url := "https://www.lrde.epita.fr/dload/spot/spot-2.12.bin"
	spot_cmd_bin := exec.Command("curl", "-L", spot_bin_url, "-o", "binary.bin")
	err = spot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spot_src_url := "https://www.lrde.epita.fr/dload/spot/spot-2.12.src.tar.gz"
	spot_cmd_src := exec.Command("curl", "-L", spot_src_url, "-o", "source.tar.gz")
	err = spot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spot_cmd_direct := exec.Command("./binary")
	err = spot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
