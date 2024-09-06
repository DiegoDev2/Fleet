package main

// Bnd - Swiss Army Knife for OSGi bundles
// Homepage: https://bnd.bndtools.org/

import (
	"fmt"
	
	"os/exec"
)

func installBnd() {
	// Método 1: Descargar y extraer .tar.gz
	bnd_tar_url := "https://search.maven.org/remotecontent?filepath=biz/aQute/bnd/biz.aQute.bnd/7.0.0/biz.aQute.bnd-7.0.0.jar"
	bnd_cmd_tar := exec.Command("curl", "-L", bnd_tar_url, "-o", "package.tar.gz")
	err := bnd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bnd_zip_url := "https://search.maven.org/remotecontent?filepath=biz/aQute/bnd/biz.aQute.bnd/7.0.0/biz.aQute.bnd-7.0.0.jar"
	bnd_cmd_zip := exec.Command("curl", "-L", bnd_zip_url, "-o", "package.zip")
	err = bnd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bnd_bin_url := "https://search.maven.org/remotecontent?filepath=biz/aQute/bnd/biz.aQute.bnd/7.0.0/biz.aQute.bnd-7.0.0.jar"
	bnd_cmd_bin := exec.Command("curl", "-L", bnd_bin_url, "-o", "binary.bin")
	err = bnd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bnd_src_url := "https://search.maven.org/remotecontent?filepath=biz/aQute/bnd/biz.aQute.bnd/7.0.0/biz.aQute.bnd-7.0.0.jar"
	bnd_cmd_src := exec.Command("curl", "-L", bnd_src_url, "-o", "source.tar.gz")
	err = bnd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bnd_cmd_direct := exec.Command("./binary")
	err = bnd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
