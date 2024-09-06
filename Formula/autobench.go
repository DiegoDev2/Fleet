package main

// Autobench - Automatic webserver benchmark tool
// Homepage: http://www.xenoclast.org/autobench/

import (
	"fmt"
	
	"os/exec"
)

func installAutobench() {
	// Método 1: Descargar y extraer .tar.gz
	autobench_tar_url := "http://www.xenoclast.org/autobench/downloads/autobench-2.1.2.tar.gz"
	autobench_cmd_tar := exec.Command("curl", "-L", autobench_tar_url, "-o", "package.tar.gz")
	err := autobench_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autobench_zip_url := "http://www.xenoclast.org/autobench/downloads/autobench-2.1.2.zip"
	autobench_cmd_zip := exec.Command("curl", "-L", autobench_zip_url, "-o", "package.zip")
	err = autobench_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autobench_bin_url := "http://www.xenoclast.org/autobench/downloads/autobench-2.1.2.bin"
	autobench_cmd_bin := exec.Command("curl", "-L", autobench_bin_url, "-o", "binary.bin")
	err = autobench_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autobench_src_url := "http://www.xenoclast.org/autobench/downloads/autobench-2.1.2.src.tar.gz"
	autobench_cmd_src := exec.Command("curl", "-L", autobench_src_url, "-o", "source.tar.gz")
	err = autobench_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autobench_cmd_direct := exec.Command("./binary")
	err = autobench_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: httperf")
	exec.Command("latte", "install", "httperf").Run()
}
