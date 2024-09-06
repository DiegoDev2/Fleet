package main

// Spotbugs - Tool for Java static analysis (FindBugs's successor)
// Homepage: https://spotbugs.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installSpotbugs() {
	// Método 1: Descargar y extraer .tar.gz
	spotbugs_tar_url := "https://repo.maven.apache.org/maven2/com/github/spotbugs/spotbugs/4.8.6/spotbugs-4.8.6.tgz"
	spotbugs_cmd_tar := exec.Command("curl", "-L", spotbugs_tar_url, "-o", "package.tar.gz")
	err := spotbugs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spotbugs_zip_url := "https://repo.maven.apache.org/maven2/com/github/spotbugs/spotbugs/4.8.6/spotbugs-4.8.6.tgz"
	spotbugs_cmd_zip := exec.Command("curl", "-L", spotbugs_zip_url, "-o", "package.zip")
	err = spotbugs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spotbugs_bin_url := "https://repo.maven.apache.org/maven2/com/github/spotbugs/spotbugs/4.8.6/spotbugs-4.8.6.tgz"
	spotbugs_cmd_bin := exec.Command("curl", "-L", spotbugs_bin_url, "-o", "binary.bin")
	err = spotbugs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spotbugs_src_url := "https://repo.maven.apache.org/maven2/com/github/spotbugs/spotbugs/4.8.6/spotbugs-4.8.6.tgz"
	spotbugs_cmd_src := exec.Command("curl", "-L", spotbugs_src_url, "-o", "source.tar.gz")
	err = spotbugs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spotbugs_cmd_direct := exec.Command("./binary")
	err = spotbugs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gradle")
exec.Command("latte", "install", "gradle")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
