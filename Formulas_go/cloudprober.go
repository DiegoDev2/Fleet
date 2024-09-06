package main

// Cloudprober - Active monitoring software to detect failures before your customers do
// Homepage: https://cloudprober.org

import (
	"fmt"
	
	"os/exec"
)

func installCloudprober() {
	// Método 1: Descargar y extraer .tar.gz
	cloudprober_tar_url := "https://github.com/cloudprober/cloudprober/archive/refs/tags/v0.13.7.tar.gz"
	cloudprober_cmd_tar := exec.Command("curl", "-L", cloudprober_tar_url, "-o", "package.tar.gz")
	err := cloudprober_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloudprober_zip_url := "https://github.com/cloudprober/cloudprober/archive/refs/tags/v0.13.7.zip"
	cloudprober_cmd_zip := exec.Command("curl", "-L", cloudprober_zip_url, "-o", "package.zip")
	err = cloudprober_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloudprober_bin_url := "https://github.com/cloudprober/cloudprober/archive/refs/tags/v0.13.7.bin"
	cloudprober_cmd_bin := exec.Command("curl", "-L", cloudprober_bin_url, "-o", "binary.bin")
	err = cloudprober_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloudprober_src_url := "https://github.com/cloudprober/cloudprober/archive/refs/tags/v0.13.7.src.tar.gz"
	cloudprober_cmd_src := exec.Command("curl", "-L", cloudprober_src_url, "-o", "source.tar.gz")
	err = cloudprober_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloudprober_cmd_direct := exec.Command("./binary")
	err = cloudprober_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
