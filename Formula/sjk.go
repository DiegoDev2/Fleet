package main

// Sjk - Swiss Java Knife
// Homepage: https://github.com/aragozin/jvm-tools

import (
	"fmt"
	
	"os/exec"
)

func installSjk() {
	// Método 1: Descargar y extraer .tar.gz
	sjk_tar_url := "https://search.maven.org/remotecontent?filepath=org/gridkit/jvmtool/sjk-plus/0.23/sjk-plus-0.23.jar"
	sjk_cmd_tar := exec.Command("curl", "-L", sjk_tar_url, "-o", "package.tar.gz")
	err := sjk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sjk_zip_url := "https://search.maven.org/remotecontent?filepath=org/gridkit/jvmtool/sjk-plus/0.23/sjk-plus-0.23.jar"
	sjk_cmd_zip := exec.Command("curl", "-L", sjk_zip_url, "-o", "package.zip")
	err = sjk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sjk_bin_url := "https://search.maven.org/remotecontent?filepath=org/gridkit/jvmtool/sjk-plus/0.23/sjk-plus-0.23.jar"
	sjk_cmd_bin := exec.Command("curl", "-L", sjk_bin_url, "-o", "binary.bin")
	err = sjk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sjk_src_url := "https://search.maven.org/remotecontent?filepath=org/gridkit/jvmtool/sjk-plus/0.23/sjk-plus-0.23.jar"
	sjk_cmd_src := exec.Command("curl", "-L", sjk_src_url, "-o", "source.tar.gz")
	err = sjk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sjk_cmd_direct := exec.Command("./binary")
	err = sjk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
