package main

// Metabase - Business intelligence report server
// Homepage: https://www.metabase.com/

import (
	"fmt"
	
	"os/exec"
)

func installMetabase() {
	// Método 1: Descargar y extraer .tar.gz
	metabase_tar_url := "https://downloads.metabase.com/v0.50.24/metabase.jar"
	metabase_cmd_tar := exec.Command("curl", "-L", metabase_tar_url, "-o", "package.tar.gz")
	err := metabase_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	metabase_zip_url := "https://downloads.metabase.com/v0.50.24/metabase.jar"
	metabase_cmd_zip := exec.Command("curl", "-L", metabase_zip_url, "-o", "package.zip")
	err = metabase_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	metabase_bin_url := "https://downloads.metabase.com/v0.50.24/metabase.jar"
	metabase_cmd_bin := exec.Command("curl", "-L", metabase_bin_url, "-o", "binary.bin")
	err = metabase_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	metabase_src_url := "https://downloads.metabase.com/v0.50.24/metabase.jar"
	metabase_cmd_src := exec.Command("curl", "-L", metabase_src_url, "-o", "source.tar.gz")
	err = metabase_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	metabase_cmd_direct := exec.Command("./binary")
	err = metabase_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: leiningen")
exec.Command("latte", "install", "leiningen")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: yarn")
exec.Command("latte", "install", "yarn")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
