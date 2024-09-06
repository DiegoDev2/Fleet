package main

// Gittuf - Security layer for Git repositories
// Homepage: https://gittuf.dev/

import (
	"fmt"
	
	"os/exec"
)

func installGittuf() {
	// Método 1: Descargar y extraer .tar.gz
	gittuf_tar_url := "https://github.com/gittuf/gittuf/archive/refs/tags/v0.5.2.tar.gz"
	gittuf_cmd_tar := exec.Command("curl", "-L", gittuf_tar_url, "-o", "package.tar.gz")
	err := gittuf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gittuf_zip_url := "https://github.com/gittuf/gittuf/archive/refs/tags/v0.5.2.zip"
	gittuf_cmd_zip := exec.Command("curl", "-L", gittuf_zip_url, "-o", "package.zip")
	err = gittuf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gittuf_bin_url := "https://github.com/gittuf/gittuf/archive/refs/tags/v0.5.2.bin"
	gittuf_cmd_bin := exec.Command("curl", "-L", gittuf_bin_url, "-o", "binary.bin")
	err = gittuf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gittuf_src_url := "https://github.com/gittuf/gittuf/archive/refs/tags/v0.5.2.src.tar.gz"
	gittuf_cmd_src := exec.Command("curl", "-L", gittuf_src_url, "-o", "source.tar.gz")
	err = gittuf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gittuf_cmd_direct := exec.Command("./binary")
	err = gittuf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
