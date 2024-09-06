package main

// PerconaToolkit - Command-line tools for MySQL, MariaDB and system tasks
// Homepage: https://www.percona.com/software/percona-toolkit/

import (
	"fmt"
	
	"os/exec"
)

func installPerconaToolkit() {
	// Método 1: Descargar y extraer .tar.gz
	perconatoolkit_tar_url := "https://www.percona.com/downloads/percona-toolkit/3.6.0/source/tarball/percona-toolkit-3.6.0.tar.gz"
	perconatoolkit_cmd_tar := exec.Command("curl", "-L", perconatoolkit_tar_url, "-o", "package.tar.gz")
	err := perconatoolkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	perconatoolkit_zip_url := "https://www.percona.com/downloads/percona-toolkit/3.6.0/source/tarball/percona-toolkit-3.6.0.zip"
	perconatoolkit_cmd_zip := exec.Command("curl", "-L", perconatoolkit_zip_url, "-o", "package.zip")
	err = perconatoolkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	perconatoolkit_bin_url := "https://www.percona.com/downloads/percona-toolkit/3.6.0/source/tarball/percona-toolkit-3.6.0.bin"
	perconatoolkit_cmd_bin := exec.Command("curl", "-L", perconatoolkit_bin_url, "-o", "binary.bin")
	err = perconatoolkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	perconatoolkit_src_url := "https://www.percona.com/downloads/percona-toolkit/3.6.0/source/tarball/percona-toolkit-3.6.0.src.tar.gz"
	perconatoolkit_cmd_src := exec.Command("curl", "-L", perconatoolkit_src_url, "-o", "source.tar.gz")
	err = perconatoolkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	perconatoolkit_cmd_direct := exec.Command("./binary")
	err = perconatoolkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: mysql-client")
exec.Command("latte", "install", "mysql-client")
}
