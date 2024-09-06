package main

// Datafusion - Apache Arrow DataFusion and Ballista query engines
// Homepage: https://arrow.apache.org/datafusion

import (
	"fmt"
	
	"os/exec"
)

func installDatafusion() {
	// Método 1: Descargar y extraer .tar.gz
	datafusion_tar_url := "https://github.com/apache/arrow-datafusion/archive/refs/tags/41.0.0.tar.gz"
	datafusion_cmd_tar := exec.Command("curl", "-L", datafusion_tar_url, "-o", "package.tar.gz")
	err := datafusion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	datafusion_zip_url := "https://github.com/apache/arrow-datafusion/archive/refs/tags/41.0.0.zip"
	datafusion_cmd_zip := exec.Command("curl", "-L", datafusion_zip_url, "-o", "package.zip")
	err = datafusion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	datafusion_bin_url := "https://github.com/apache/arrow-datafusion/archive/refs/tags/41.0.0.bin"
	datafusion_cmd_bin := exec.Command("curl", "-L", datafusion_bin_url, "-o", "binary.bin")
	err = datafusion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	datafusion_src_url := "https://github.com/apache/arrow-datafusion/archive/refs/tags/41.0.0.src.tar.gz"
	datafusion_cmd_src := exec.Command("curl", "-L", datafusion_src_url, "-o", "source.tar.gz")
	err = datafusion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	datafusion_cmd_direct := exec.Command("./binary")
	err = datafusion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
