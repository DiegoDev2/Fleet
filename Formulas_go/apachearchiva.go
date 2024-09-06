package main

// ApacheArchiva - Build Artifact Repository Manager
// Homepage: https://archiva.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installApacheArchiva() {
	// Método 1: Descargar y extraer .tar.gz
	apachearchiva_tar_url := "https://www.apache.org/dyn/closer.lua?path=archiva/2.2.10/binaries/apache-archiva-2.2.10-bin.tar.gz"
	apachearchiva_cmd_tar := exec.Command("curl", "-L", apachearchiva_tar_url, "-o", "package.tar.gz")
	err := apachearchiva_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apachearchiva_zip_url := "https://www.apache.org/dyn/closer.lua?path=archiva/2.2.10/binaries/apache-archiva-2.2.10-bin.zip"
	apachearchiva_cmd_zip := exec.Command("curl", "-L", apachearchiva_zip_url, "-o", "package.zip")
	err = apachearchiva_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apachearchiva_bin_url := "https://www.apache.org/dyn/closer.lua?path=archiva/2.2.10/binaries/apache-archiva-2.2.10-bin.bin"
	apachearchiva_cmd_bin := exec.Command("curl", "-L", apachearchiva_bin_url, "-o", "binary.bin")
	err = apachearchiva_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apachearchiva_src_url := "https://www.apache.org/dyn/closer.lua?path=archiva/2.2.10/binaries/apache-archiva-2.2.10-bin.src.tar.gz"
	apachearchiva_cmd_src := exec.Command("curl", "-L", apachearchiva_src_url, "-o", "source.tar.gz")
	err = apachearchiva_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apachearchiva_cmd_direct := exec.Command("./binary")
	err = apachearchiva_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ant")
exec.Command("latte", "install", "ant")
	fmt.Println("Instalando dependencia: java-service-wrapper")
exec.Command("latte", "install", "java-service-wrapper")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
