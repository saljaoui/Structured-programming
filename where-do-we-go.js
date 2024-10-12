import { places } from "./where-do-we-go.data.js";

export const explore = () => {
    // Sort places from north to south
    const sortedPlaces = places.sort((a, b) => b.coordinates[0] - a.coordinates[0]);
    // Create and append sections for each place
    sortedPlaces.forEach(place => {
      const section = document.createElement('section');
      const firstPart = place.name.split(",")[0].trim();
      section.style.backgroundImage = `url('${firstPart.toLowerCase().replace(/\s+/g, '-')}.jpg')`
;
      section.style.backgroundSize = 'cover';
      section.style.backgroundPosition = 'center';
      section.style.height = '100vh';
      section.style.width = '100vw';
      section.style.display = 'flex';
      section.style.justifyContent = 'center';
      section.style.alignItems = 'center';
  
      const nameElement = document.createElement('h2');
      nameElement.textContent = place.name;
      nameElement.style.color = 'white';
      nameElement.style.fontSize = '3rem';
      nameElement.style.textShadow = '2px 2px 4px rgba(0,0,0,0.5)';
  
      section.appendChild(nameElement);
      document.body.appendChild(section);
    });
  };