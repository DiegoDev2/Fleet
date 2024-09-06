package main

// Metis - Programs that partition graphs and order matrices
// Homepage: http://glaros.dtc.umn.edu/gkhome/views/metis

import (
	"fmt"
	
	"os/exec"
)

func installMetis() {
	// Método 1: Descargar y extraer .tar.gz
	metis_tar_url := "http://glaros.dtc.umn.edu/gkhome/fetch/sw/metis/metis-5.1.0.tar.gz"
	metis_cmd_tar := exec.Command("curl", "-L", metis_tar_url, "-o", "package.tar.gz")
	err := metis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	metis_zip_url := "http://glaros.dtc.umn.edu/gkhome/fetch/sw/metis/metis-5.1.0.zip"
	metis_cmd_zip := exec.Command("curl", "-L", metis_zip_url, "-o", "package.zip")
	err = metis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	metis_bin_url := "http://glaros.dtc.umn.edu/gkhome/fetch/sw/metis/metis-5.1.0.bin"
	metis_cmd_bin := exec.Command("curl", "-L", metis_bin_url, "-o", "binary.bin")
	err = metis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	metis_src_url := "http://glaros.dtc.umn.edu/gkhome/fetch/sw/metis/metis-5.1.0.src.tar.gz"
	metis_cmd_src := exec.Command("curl", "-L", metis_src_url, "-o", "source.tar.gz")
	err = metis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	metis_cmd_direct := exec.Command("./binary")
	err = metis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
