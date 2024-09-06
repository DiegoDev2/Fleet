package main

// FdkAac - Standalone library of the Fraunhofer FDK AAC code from Android
// Homepage: https://sourceforge.net/projects/opencore-amr/

import (
	"fmt"
	
	"os/exec"
)

func installFdkAac() {
	// Método 1: Descargar y extraer .tar.gz
	fdkaac_tar_url := "https://downloads.sourceforge.net/project/opencore-amr/fdk-aac/fdk-aac-2.0.3.tar.gz"
	fdkaac_cmd_tar := exec.Command("curl", "-L", fdkaac_tar_url, "-o", "package.tar.gz")
	err := fdkaac_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fdkaac_zip_url := "https://downloads.sourceforge.net/project/opencore-amr/fdk-aac/fdk-aac-2.0.3.zip"
	fdkaac_cmd_zip := exec.Command("curl", "-L", fdkaac_zip_url, "-o", "package.zip")
	err = fdkaac_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fdkaac_bin_url := "https://downloads.sourceforge.net/project/opencore-amr/fdk-aac/fdk-aac-2.0.3.bin"
	fdkaac_cmd_bin := exec.Command("curl", "-L", fdkaac_bin_url, "-o", "binary.bin")
	err = fdkaac_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fdkaac_src_url := "https://downloads.sourceforge.net/project/opencore-amr/fdk-aac/fdk-aac-2.0.3.src.tar.gz"
	fdkaac_cmd_src := exec.Command("curl", "-L", fdkaac_src_url, "-o", "source.tar.gz")
	err = fdkaac_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fdkaac_cmd_direct := exec.Command("./binary")
	err = fdkaac_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
