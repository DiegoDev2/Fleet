package main

// Soci - Database access library for C++
// Homepage: https://soci.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSoci() {
	// Método 1: Descargar y extraer .tar.gz
	soci_tar_url := "https://downloads.sourceforge.net/project/soci/soci/soci-4.0.3/soci-4.0.3.zip"
	soci_cmd_tar := exec.Command("curl", "-L", soci_tar_url, "-o", "package.tar.gz")
	err := soci_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	soci_zip_url := "https://downloads.sourceforge.net/project/soci/soci/soci-4.0.3/soci-4.0.3.zip"
	soci_cmd_zip := exec.Command("curl", "-L", soci_zip_url, "-o", "package.zip")
	err = soci_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	soci_bin_url := "https://downloads.sourceforge.net/project/soci/soci/soci-4.0.3/soci-4.0.3.zip"
	soci_cmd_bin := exec.Command("curl", "-L", soci_bin_url, "-o", "binary.bin")
	err = soci_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	soci_src_url := "https://downloads.sourceforge.net/project/soci/soci/soci-4.0.3/soci-4.0.3.zip"
	soci_cmd_src := exec.Command("curl", "-L", soci_src_url, "-o", "source.tar.gz")
	err = soci_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	soci_cmd_direct := exec.Command("./binary")
	err = soci_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
}
