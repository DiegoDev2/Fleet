package main

// Aggregate - Optimizes lists of prefixes to reduce list lengths
// Homepage: https://web.archive.org/web/20160716192438/freecode.com/projects/aggregate/

import (
	"fmt"
	
	"os/exec"
)

func installAggregate() {
	// Método 1: Descargar y extraer .tar.gz
	aggregate_tar_url := "https://ftp.isc.org/isc/aggregate/aggregate-1.6.tar.gz"
	aggregate_cmd_tar := exec.Command("curl", "-L", aggregate_tar_url, "-o", "package.tar.gz")
	err := aggregate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aggregate_zip_url := "https://ftp.isc.org/isc/aggregate/aggregate-1.6.zip"
	aggregate_cmd_zip := exec.Command("curl", "-L", aggregate_zip_url, "-o", "package.zip")
	err = aggregate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aggregate_bin_url := "https://ftp.isc.org/isc/aggregate/aggregate-1.6.bin"
	aggregate_cmd_bin := exec.Command("curl", "-L", aggregate_bin_url, "-o", "binary.bin")
	err = aggregate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aggregate_src_url := "https://ftp.isc.org/isc/aggregate/aggregate-1.6.src.tar.gz"
	aggregate_cmd_src := exec.Command("curl", "-L", aggregate_src_url, "-o", "source.tar.gz")
	err = aggregate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aggregate_cmd_direct := exec.Command("./binary")
	err = aggregate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
