package main

// Qstat - Query Quake servers from the command-line
// Homepage: https://github.com/Unity-Technologies/qstat

import (
	"fmt"
	
	"os/exec"
)

func installQstat() {
	// Método 1: Descargar y extraer .tar.gz
	qstat_tar_url := "https://github.com/Unity-Technologies/qstat/archive/refs/tags/v2.17.tar.gz"
	qstat_cmd_tar := exec.Command("curl", "-L", qstat_tar_url, "-o", "package.tar.gz")
	err := qstat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qstat_zip_url := "https://github.com/Unity-Technologies/qstat/archive/refs/tags/v2.17.zip"
	qstat_cmd_zip := exec.Command("curl", "-L", qstat_zip_url, "-o", "package.zip")
	err = qstat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qstat_bin_url := "https://github.com/Unity-Technologies/qstat/archive/refs/tags/v2.17.bin"
	qstat_cmd_bin := exec.Command("curl", "-L", qstat_bin_url, "-o", "binary.bin")
	err = qstat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qstat_src_url := "https://github.com/Unity-Technologies/qstat/archive/refs/tags/v2.17.src.tar.gz"
	qstat_cmd_src := exec.Command("curl", "-L", qstat_src_url, "-o", "source.tar.gz")
	err = qstat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qstat_cmd_direct := exec.Command("./binary")
	err = qstat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
