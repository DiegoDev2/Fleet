package main

// Scalariform - Scala source code formatter
// Homepage: https://github.com/scala-ide/scalariform

import (
	"fmt"
	
	"os/exec"
)

func installScalariform() {
	// Método 1: Descargar y extraer .tar.gz
	scalariform_tar_url := "https://github.com/scala-ide/scalariform/releases/download/0.2.10/scalariform.jar"
	scalariform_cmd_tar := exec.Command("curl", "-L", scalariform_tar_url, "-o", "package.tar.gz")
	err := scalariform_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scalariform_zip_url := "https://github.com/scala-ide/scalariform/releases/download/0.2.10/scalariform.jar"
	scalariform_cmd_zip := exec.Command("curl", "-L", scalariform_zip_url, "-o", "package.zip")
	err = scalariform_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scalariform_bin_url := "https://github.com/scala-ide/scalariform/releases/download/0.2.10/scalariform.jar"
	scalariform_cmd_bin := exec.Command("curl", "-L", scalariform_bin_url, "-o", "binary.bin")
	err = scalariform_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scalariform_src_url := "https://github.com/scala-ide/scalariform/releases/download/0.2.10/scalariform.jar"
	scalariform_cmd_src := exec.Command("curl", "-L", scalariform_src_url, "-o", "source.tar.gz")
	err = scalariform_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scalariform_cmd_direct := exec.Command("./binary")
	err = scalariform_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sbt")
exec.Command("latte", "install", "sbt")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
