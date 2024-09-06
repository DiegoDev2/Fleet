package main

// Mmdbinspect - Look up records for one or more IPs/networks in one or more .mmdb databases
// Homepage: https://github.com/maxmind/mmdbinspect

import (
	"fmt"
	
	"os/exec"
)

func installMmdbinspect() {
	// Método 1: Descargar y extraer .tar.gz
	mmdbinspect_tar_url := "https://github.com/maxmind/mmdbinspect/archive/refs/tags/v0.2.0.tar.gz"
	mmdbinspect_cmd_tar := exec.Command("curl", "-L", mmdbinspect_tar_url, "-o", "package.tar.gz")
	err := mmdbinspect_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mmdbinspect_zip_url := "https://github.com/maxmind/mmdbinspect/archive/refs/tags/v0.2.0.zip"
	mmdbinspect_cmd_zip := exec.Command("curl", "-L", mmdbinspect_zip_url, "-o", "package.zip")
	err = mmdbinspect_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mmdbinspect_bin_url := "https://github.com/maxmind/mmdbinspect/archive/refs/tags/v0.2.0.bin"
	mmdbinspect_cmd_bin := exec.Command("curl", "-L", mmdbinspect_bin_url, "-o", "binary.bin")
	err = mmdbinspect_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mmdbinspect_src_url := "https://github.com/maxmind/mmdbinspect/archive/refs/tags/v0.2.0.src.tar.gz"
	mmdbinspect_cmd_src := exec.Command("curl", "-L", mmdbinspect_src_url, "-o", "source.tar.gz")
	err = mmdbinspect_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mmdbinspect_cmd_direct := exec.Command("./binary")
	err = mmdbinspect_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
