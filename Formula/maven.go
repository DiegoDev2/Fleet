package main

// Maven - Java-based project management
// Homepage: https://maven.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installMaven() {
	// Método 1: Descargar y extraer .tar.gz
	maven_tar_url := "https://www.apache.org/dyn/closer.lua?path=maven/maven-3/3.9.9/binaries/apache-maven-3.9.9-bin.tar.gz"
	maven_cmd_tar := exec.Command("curl", "-L", maven_tar_url, "-o", "package.tar.gz")
	err := maven_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	maven_zip_url := "https://www.apache.org/dyn/closer.lua?path=maven/maven-3/3.9.9/binaries/apache-maven-3.9.9-bin.zip"
	maven_cmd_zip := exec.Command("curl", "-L", maven_zip_url, "-o", "package.zip")
	err = maven_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	maven_bin_url := "https://www.apache.org/dyn/closer.lua?path=maven/maven-3/3.9.9/binaries/apache-maven-3.9.9-bin.bin"
	maven_cmd_bin := exec.Command("curl", "-L", maven_bin_url, "-o", "binary.bin")
	err = maven_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	maven_src_url := "https://www.apache.org/dyn/closer.lua?path=maven/maven-3/3.9.9/binaries/apache-maven-3.9.9-bin.src.tar.gz"
	maven_cmd_src := exec.Command("curl", "-L", maven_src_url, "-o", "source.tar.gz")
	err = maven_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	maven_cmd_direct := exec.Command("./binary")
	err = maven_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
